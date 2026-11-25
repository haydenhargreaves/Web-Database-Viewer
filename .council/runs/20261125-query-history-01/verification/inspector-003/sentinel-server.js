'use strict';
const fs = require('fs'), path = require('path'), vm = require('vm');
const root = process.cwd();
let source = fs.readFileSync(path.join(root, 'web/tests/query-history-server.js'), 'utf8');
source = source.replace("const root = path.resolve(__dirname, '../..');", "const root = process.cwd();")
 .replace('web/tests/query-history.html', '.council/runs/20261125-query-history-01/verification/inspector-003/sentinel.html')
 .replace('web/tests/query-history.js', '.council/runs/20261125-query-history-01/verification/inspector-003/sentinel.js')
 .replace("res.end('result')", "res.end('synthetic-result')");
new Function('require', 'process', source)(require, process);
