package query

// PostgreSQL queries
// Drivers: postgres
const GET_TABLE_LIST_PSQL string = `SELECT table_name FROM information_schema.tables WHERE table_schema = 'public' AND table_type = 'BASE TABLE';`
const GET_VIEW_LIST_PSQL string = `SELECT table_name FROM information_schema.views WHERE table_schema = 'public';`
const GET_TABLE_PK_PSQL string = `SELECT kcu.column_name FROM information_schema.table_constraints tc JOIN  information_schema.key_column_usage kcu  ON tc.constraint_name = kcu.constraint_name AND tc.table_schema = kcu.table_schema WHERE tc.constraint_type = 'PRIMARY KEY' AND tc.table_name = '%s';`
const GET_TABLE_FKS_PSQL string = `SELECT tc.table_schema, tc.table_name, kcu.column_name, ccu.table_schema AS foreign_table_schema, ccu.table_name AS foreign_table_name, ccu.column_name AS foreign_column_name FROM information_schema.table_constraints AS tc JOIN information_schema.key_column_usage AS kcu ON tc.constraint_name = kcu.constraint_name AND tc.table_schema = kcu.table_schema JOIN information_schema.constraint_column_usage AS ccu ON ccu.constraint_name = tc.constraint_name AND ccu.table_schema = tc.table_schema WHERE tc.constraint_type = 'FOREIGN KEY' AND tc.table_name = '%s';`
const GET_TABLE_RESTRAINS_PSQL string = `SELECT c.column_name, c.is_nullable, c.data_type, c.character_maximum_length, t.typname AS enum_type FROM information_schema.columns c JOIN pg_type t ON c.udt_name = t.typname WHERE c.table_name = '%s';`
const GET_TABLE_UNIQUE_COLS_PSQL string = `SELECT kcu.column_name FROM information_schema.table_constraints AS tc JOIN information_schema.key_column_usage AS kcu ON tc.constraint_name = kcu.constraint_name AND tc.table_schema = kcu.table_schema WHERE tc.constraint_type = 'UNIQUE' AND kcu.table_name = '%s';`

const GET_ENUM_LIST_PSQL string = `SELECT t.typname AS enum_name, e.enumlabel AS enum_value FROM pg_type t JOIN pg_enum e ON t.oid = e.enumtypid JOIN pg_namespace n ON n.oid = t.typnamespace WHERE t.typcategory = 'E' AND n.nspname NOT IN ('pg_catalog', 'information_schema')  ORDER BY t.typname, e.enumsortorder;`
const GET_ROUTINE_LIST_PSQL string = `SELECT p.oid::text, n.nspname, p.proname, p.oid::regprocedure::text, CASE p.prokind WHEN 'p' THEN 'procedure' ELSE 'function' END FROM pg_proc p JOIN pg_namespace n ON n.oid = p.pronamespace WHERE p.prokind IN ('f', 'p') AND n.nspname NOT IN ('pg_catalog', 'information_schema') ORDER BY n.nspname, p.proname, p.oid::regprocedure::text;`
const GET_ROUTINE_DEFINITION_PSQL string = `SELECT pg_get_functiondef(%d);`

// MySQL queries
// Drivers: mysql, mariadb
const GET_TABLE_LIST_MYSQL string = `SELECT table_name FROM information_schema.tables WHERE table_schema = DATABASE();`
const GET_VIEW_LIST_MYSQL string = `SELECT table_name FROM information_schema.tables WHERE table_schema = DATABASE() AND table_type = 'VIEW';`
const GET_TABLE_PK_MYSQL string = `SELECT column_name FROM information_schema.key_column_usage WHERE table_schema = DATABASE() AND table_name = '%s' AND constraint_name = 'PRIMARY';`
const GET_TABLE_FKS_MYSQL string = `SELECT tc.constraint_schema AS table_schema, tc.table_name, kcu.column_name, kcu.referenced_table_schema AS foreign_table_schema, kcu.referenced_table_name AS foreign_table_name, kcu.referenced_column_name AS foreign_column_name FROM information_schema.table_constraints AS tc JOIN information_schema.key_column_usage AS kcu ON tc.constraint_name = kcu.constraint_name AND tc.constraint_schema = kcu.constraint_schema WHERE tc.constraint_type = 'FOREIGN KEY' AND tc.table_name = '%s';`
const GET_TABLE_RESTRAINS_MYSQL string = `SELECT c.column_name, c.is_nullable, c.data_type, CAST(c.character_maximum_length AS UNSIGNED) AS max_length, CASE WHEN c.data_type = 'enum' THEN SUBSTRING(c.column_type, 6, LENGTH(c.column_type) - 6 - 1) ELSE NULL END AS enum_type FROM information_schema.columns c WHERE c.table_name = '%s' AND c.table_schema = DATABASE();`
const GET_TABLE_UNIQUE_COLS_MYSQL string = `SELECT kcu.column_name FROM information_schema.table_constraints AS tc JOIN information_schema.key_column_usage AS kcu ON tc.constraint_name = kcu.constraint_name AND tc.constraint_schema = kcu.table_schema WHERE tc.constraint_type = 'UNIQUE' AND kcu.table_name = '%s' AND kcu.table_schema = DATABASE();`
const GET_ROUTINE_LIST_MYSQL string = `SELECT routine_name, routine_schema, routine_name, routine_name, LOWER(routine_type) FROM information_schema.routines WHERE routine_schema = DATABASE() ORDER BY routine_type, routine_name;`

// Sqlite queries
// Drivers: sqlite3
const GET_TABLE_LIST_SQLITE string = `SELECT name FROM sqlite_master WHERE type = 'table' AND name NOT LIKE 'sqlite_%';`
const GET_VIEW_LIST_SQLITE string = `SELECT name FROM sqlite_master WHERE type = 'view';`
const GET_TABLE_PK_SQLITE string = `SELECT name FROM pragma_table_info('%s') WHERE pk > 0;`
const GET_TABLE_FKS_SQLITE string = `SELECT NULL as table_schema, NULL as table_name, "from" as column_name, NULL as foreign_table_schema, "table" as foreign_table_name, "to" as foreign_column_name FROM pragma_foreign_key_list('%s');`
const GET_TABLE_RESTRAINS_SQLITE string = `SELECT name AS column_name, CASE WHEN "notnull" = 0 THEN  'YES' ELSE 'NO' END as is_nullable, type AS data_type, NULL as character_max_length, NULL as enum_type FROM pragma_table_info('%s');`
const GET_TABLE_UNIQUE_COLS_SQLITE string = `WITH idx AS (SELECT name FROM pragma_index_list('%s') WHERE [unique] = 1) SELECT DISTINCT ii.name as column_name FROM idx il JOIN pragma_index_info(il.name) ii;`

// SQL Server queries
// Drivers: sqlserver
const GET_TABLE_LIST_MSSQL string = `SELECT TABLE_NAME FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_TYPE = 'BASE TABLE' AND TABLE_SCHEMA NOT IN ('sys', 'INFORMATION_SCHEMA');`
const GET_VIEW_LIST_MSSQL string = `SELECT TABLE_NAME FROM INFORMATION_SCHEMA.VIEWS WHERE TABLE_SCHEMA NOT IN ('sys', 'INFORMATION_SCHEMA');`
const GET_TABLE_PK_MSSQL string = `SELECT kcu.COLUMN_NAME FROM INFORMATION_SCHEMA.TABLE_CONSTRAINTS tc JOIN INFORMATION_SCHEMA.KEY_COLUMN_USAGE kcu ON tc.CONSTRAINT_NAME = kcu.CONSTRAINT_NAME AND tc.TABLE_SCHEMA = kcu.TABLE_SCHEMA WHERE tc.CONSTRAINT_TYPE = 'PRIMARY KEY' AND tc.TABLE_NAME = '%s';`
const GET_TABLE_FKS_MSSQL string = `SELECT OBJECT_SCHEMA_NAME(fk.parent_object_id) AS table_schema, t_parent.name AS table_name, c_parent.name AS column_name, OBJECT_SCHEMA_NAME(fk.referenced_object_id) AS foreign_table_schema, t_ref.name AS foreign_table_name, c_ref.name AS foreign_column_name FROM sys.foreign_key_columns fkc JOIN sys.foreign_keys fk ON fk.object_id = fkc.constraint_object_id JOIN sys.tables t_parent ON t_parent.object_id = fkc.parent_object_id JOIN sys.columns c_parent ON c_parent.object_id = fkc.parent_object_id AND c_parent.column_id = fkc.parent_column_id JOIN sys.tables t_ref ON t_ref.object_id = fkc.referenced_object_id JOIN sys.columns c_ref ON c_ref.object_id = fkc.referenced_object_id AND c_ref.column_id = fkc.referenced_column_id WHERE t_parent.name = '%s';`
const GET_TABLE_RESTRAINS_MSSQL string = `SELECT COLUMN_NAME, IS_NULLABLE, DATA_TYPE, CAST(CHARACTER_MAXIMUM_LENGTH AS BIGINT) AS CHARACTER_MAXIMUM_LENGTH, NULL AS enum_type FROM INFORMATION_SCHEMA.COLUMNS WHERE TABLE_NAME = '%s' ORDER BY ORDINAL_POSITION;`
const GET_TABLE_UNIQUE_COLS_MSSQL string = `SELECT kcu.COLUMN_NAME FROM INFORMATION_SCHEMA.TABLE_CONSTRAINTS tc JOIN INFORMATION_SCHEMA.KEY_COLUMN_USAGE kcu ON tc.CONSTRAINT_NAME = kcu.CONSTRAINT_NAME AND tc.TABLE_SCHEMA = kcu.TABLE_SCHEMA WHERE tc.CONSTRAINT_TYPE = 'UNIQUE' AND tc.TABLE_NAME = '%s';`
const GET_ROUTINE_LIST_MSSQL string = `SELECT o.object_id, SCHEMA_NAME(o.schema_id), o.name, CONCAT(SCHEMA_NAME(o.schema_id), '.', o.name), CASE WHEN o.type IN ('P', 'PC') THEN 'procedure' ELSE 'function' END FROM sys.objects o WHERE o.type IN ('P', 'PC', 'FN', 'TF', 'IF', 'FS', 'FT') AND o.is_ms_shipped = 0 ORDER BY SCHEMA_NAME(o.schema_id), o.name;`
const GET_ROUTINE_DEFINITION_MSSQL string = `SELECT OBJECT_DEFINITION(%d);`
