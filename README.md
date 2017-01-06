# Simple Pascal Interpreter Implemented in Go

# Grammar

* program: PROGRAM variable SEMI block DOT
* block: declarations compound_statements
* declarations: VAR (variable_declaration SEMI)+ | empty
* variable_declaration: ID (COMMA ID)* COLON type_spec
* type_spec: INTEGER | REAL
* compound_statements: BEGIN statement_list END
* statement_list: statement (SEMI statement_list)* | empty
* statement: compound_statement | assign_statement | empty
* assign_statement: variable ASSIGN expr
* empty:
* expr: term ((PLUS | MINUS) term)*
* term: factor ((MUL | DIV_INTEGER | DIV_REAL) factor)*
* factor: PLUS factor | MINUS factor | INTEGER_CONST | REAL_CONST | LPARAN expr RPARAN | variable
* variable: ID