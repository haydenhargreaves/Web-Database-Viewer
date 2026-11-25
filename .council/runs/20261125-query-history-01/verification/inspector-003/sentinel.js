'use strict';
(async () => {
 let msg;
 try {
  await new Promise(resolve => window.addEventListener('load', resolve, {once:true})); localStorage.clear();
  setTimeout(() => document.getElementById('run').click(), 50);
  await new Promise((resolve,reject)=>{document.body.addEventListener('htmx:afterRequest',resolve,{once:true}); setTimeout(()=>reject(Error('request timeout')),4000)});
  const key='query-history:v1:'+'a'.repeat(64), raw=localStorage.getItem(key), entries=JSON.parse(raw);
  if(localStorage.length!==1 || localStorage.key(0)!==key || entries.length!==1 || entries[0].sql!=='select 123' || Object.keys(entries[0]).sort().join(',')!=='sql,timestamp') throw Error('unexpected history shape');
  for(const s of ['postgres://','synthetic-user','synthetic-password','synthetic-host','synthetic-driver','synthetic-result']) if((key+raw).includes(s)) throw Error('sentinel leaked: '+s);
  if(!document.getElementById('result').textContent.includes('synthetic-result')) throw Error('result sentinel absent');
  msg='PASS synthetic URL, user, password, driver and result sentinels absent from key and stored data; sql,timestamp only';
 } catch(e) { msg='FAIL '+e.message; }
 await fetch('/done',{method:'POST',body:msg});
})();
