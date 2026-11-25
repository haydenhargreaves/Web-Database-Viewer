'use strict';
// Local synthetic server; no real databases or credentials.
const http = require('http');
const fs = require('fs');
const path = require('path');
const { spawn } = require('child_process');
const root = path.resolve(__dirname, '../..');
const evidence = process.env.QUERY_HISTORY_EVIDENCE_DIR;
if (!evidence || !path.isAbsolute(evidence)) throw new Error('QUERY_HISTORY_EVIDENCE_DIR must be an absolute directory');
let posts = 0;
const payloads = [];
let finish;
const files = {
  '/': ['web/tests/query-history.html', 'text/html'],
  '/test.js': ['web/tests/query-history.js', 'text/javascript'],
  '/query-history.js': ['web/static/scripts/query-history.js', 'text/javascript'],
  '/htmx.js': ['.council/runs/20261125-query-history-01/verification/engineer-001/htmx.min.js', 'text/javascript']
};
const server = http.createServer((req, res) => {
  if (req.method === 'POST' && req.url === '/v1/api/query') {
    posts++; let body = ''; req.on('data', b => body += b); req.on('end', () => {
      const sql = new URLSearchParams(body).get('sql');
      payloads.push(sql);
      const headers = { 'Content-Type': 'text/html' };
      if (sql !== 'missing identity') headers['X-Query-History-Connection'] = sql === 'invalid identity' ? 'invalid' : (sql === 'race B' ? 'b' : 'a').repeat(64);
      const respond = () => { res.writeHead(sql === 'SQL error' ? 500 : 200, headers); res.end('result'); };
      if (sql === 'race A') setTimeout(respond, 400); else respond();
    }); return;
  }
  if (req.url === '/counts') { res.end(String(posts)); return; }
  if (req.url === '/payloads') { res.setHeader('Content-Type', 'application/json'); res.end(JSON.stringify(payloads)); return; }
  if (req.method === 'POST' && req.url === '/done') {
    let body = ''; req.on('data', b => body += b); req.on('end', () => { res.end('ok'); finish(body); }); return;
  }
  const file = files[req.url];
  if (!file) { res.writeHead(404); res.end(); return; }
  res.writeHead(200, { 'Content-Type': file[1] }); fs.createReadStream(path.join(root, file[0])).pipe(res);
});
server.listen(0, '127.0.0.1', () => {
  const profile = fs.mkdtempSync(path.join(evidence, 'chromium-profile-'));
  const args = ['--headless', '--disable-gpu', `--user-data-dir=${profile}`, `http://127.0.0.1:${server.address().port}/`];
  const child = spawn('chromium', args, { stdio: ['ignore', 'pipe', 'pipe'] });
  let stderr = '', result;
  child.stdout.resume(); child.stderr.on('data', b => stderr += b);
  finish = body => { result = body; child.kill(); };
  const timer = setTimeout(() => child.kill(), 30000);
  child.on('error', error => { console.error(error); clearTimeout(timer); server.close(); });
  child.on('close', code => {
    clearTimeout(timer); server.close(); fs.rmSync(profile, { recursive: true, force: true });
    console.log(result || 'FAIL browser did not report a result');
    if (!result || !result.startsWith('PASS ')) { console.error(stderr.slice(-3000)); process.exitCode = 1; }
  });
});
