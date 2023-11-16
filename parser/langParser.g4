parser grammar langParser;


options {
    tokenVocab=langLexer;
}

program: var_declaration* EOF
       ;
var_declaration: IDENTIFIER IDENTIFIER ASSIGN number;
number: INTEGER
      | REAL
      ;