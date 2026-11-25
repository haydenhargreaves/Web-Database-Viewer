(() => {
  'use strict';
  const prefix = 'query-history:v1:';
  const validID = id => typeof id === 'string' && /^[0-9a-f]{64}$/.test(id);
  const requests = new WeakMap();
  let inputGeneration = 0;
  let blockedGeneration = -1;
  let lastTimestamp = 0;
  let switching = false;

  function selectedID() {
    const select = document.getElementById('connected-database');
    const id = !switching && select && select.selectedOptions[0]?.dataset.queryHistoryId;
    return validID(id) ? id : null;
  }

  function read(id) {
    try {
      const raw = localStorage.getItem(prefix + id);
      if (raw === null) return [];
      const entries = JSON.parse(raw);
      if (!Array.isArray(entries) || entries.some(e => !e || typeof e.sql !== 'string' ||
        typeof e.timestamp !== 'string' || !Number.isFinite(Date.parse(e.timestamp)))) return [];
      return entries.slice(0, 50).map(e => ({ sql: e.sql, timestamp: e.timestamp }));
    } catch (_) { return []; }
  }

  function write(id, entries) {
    try { localStorage.setItem(prefix + id, JSON.stringify(entries.slice(0, 50))); return true; }
    catch (_) { return false; }
  }

  function invalidatePending() { blockedGeneration = inputGeneration; }

  function render() {
    const list = document.getElementById('query-history-list');
    const status = document.getElementById('query-history-status');
    const clear = document.getElementById('query-history-clear');
    if (!list || !status || !clear) return;
    const id = selectedID();
    list.replaceChildren();
    clear.disabled = !id;
    if (!id) { status.textContent = 'History unavailable for this connection.'; return; }
    const entries = read(id);
    status.textContent = entries.length ? '' : 'No query history for this connection.';
    entries.forEach((entry, index) => {
      const item = document.createElement('li');
      item.className = 'mt-2 p-2 border rounded-md';
      const load = document.createElement('button');
      load.type = 'button';
      load.className = 'block w-full text-left p-2 focus:ring-2 focus:ring-blue-500';
      const label = document.createElement('span');
      label.className = 'sr-only';
      label.textContent = 'Load query: ';
      const preview = document.createElement('span');
      preview.style.cssText = 'display:block;white-space:pre-wrap;overflow-wrap:anywhere';
      preview.textContent = entry.sql;
      load.title = entry.sql;
      const time = document.createElement('time');
      time.dateTime = entry.timestamp;
      time.textContent = new Date(entry.timestamp).toLocaleString();
      load.append(label, preview, time);
      load.addEventListener('click', () => {
        invalidatePending();
        const editor = document.getElementById('sql');
        if (editor) { editor.value = entry.sql; editor.focus(); }
      });
      const remove = document.createElement('button');
      remove.type = 'button';
      remove.className = 'p-2 border rounded-md focus:ring-2 focus:ring-blue-500';
      remove.textContent = 'Remove';
      remove.setAttribute('aria-label', 'Remove query from ' + new Date(entry.timestamp).toLocaleString());
      remove.addEventListener('click', () => {
        invalidatePending();
        const current = read(id);
        // Re-read before mutation; match one occurrence rather than using a stale index.
        const matches = current.map((e, i) => e.sql === entry.sql && e.timestamp === entry.timestamp ? i : -1).filter(i => i >= 0);
        const occurrence = entries.slice(0, index + 1).filter(e => e.sql === entry.sql && e.timestamp === entry.timestamp).length - 1;
        if (matches[occurrence] !== undefined) current.splice(matches[occurrence], 1);
        write(id, current);
        render();
      });
      item.append(load, remove);
      list.append(item);
    });
  }

  document.addEventListener('input', e => {
    if (e.target.id === 'sql') inputGeneration++;
  }, true);
  document.addEventListener('change', e => {
    if (e.target.id === 'connected-database') { switching = true; render(); }
  }, true);
  document.addEventListener('click', e => {
    if (e.target.closest('#query-history-clear')) {
      invalidatePending();
      const id = selectedID();
      if (id) {
        try { localStorage.removeItem(prefix + id); } catch (_) { /* Querying remains available. */ }
        render();
      }
    }
  });
  document.body.addEventListener('htmx:beforeRequest', e => {
    try {
      const d = e.detail;
      const config = d.requestConfig;
      if (!config || config.verb.toLowerCase() !== 'post' || config.path.split('?')[0] !== '/v1/api/query') return;
      if (d.elt.id === 'sql' && blockedGeneration === inputGeneration) { e.preventDefault(); return; }
      const sql = config.parameters.sql;
      if (typeof sql !== 'string' || !sql.trim()) return;
      lastTimestamp = Math.max(Date.now(), lastTimestamp + 1);
      requests.set(d.xhr, { sql, timestamp: new Date(lastTimestamp).toISOString() });
    } catch (_) { /* History must never block query execution. */ }
  });
  document.body.addEventListener('htmx:afterRequest', e => {
    try {
      const xhr = e.detail.xhr;
      const entry = requests.get(xhr);
      if (entry) {
        requests.delete(xhr);
        const id = xhr.status && xhr.getResponseHeader('X-Query-History-Connection');
        if (validID(id)) {
          const entries = read(id);
          entries.push(entry);
          entries.sort((a, b) => Date.parse(b.timestamp) - Date.parse(a.timestamp));
          write(id, entries);
          render();
        }
      }
      if (e.detail.target?.id === 'connected-database' && e.detail.failed) { switching = false; render(); }
    } catch (_) { /* Storage and unavailable responses are optional. */ }
  });
  document.body.addEventListener('htmx:afterSwap', e => {
    if (e.detail.target?.id === 'connected-database' || e.detail.elt?.id === 'connected-database') {
      switching = false;
      render();
    }
  });
  window.addEventListener('storage', e => { if (!e.key || e.key.startsWith(prefix)) render(); });
  render();
})();
