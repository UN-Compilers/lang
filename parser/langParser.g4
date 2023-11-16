parser grammar langParser;


options {
    tokenVocab=langLexer;
}

program: var_declaration* EOF
       ;
var_declaration: IDENTIFIER IDENTIFIER ASSIGN exp;
number: INTEGER
      | REAL
      ;

exp: exp TIMES exp
   | exp PLUS exp
   | LPAR exp RPAR
   | atom
   ;

 atom: number
     | IDENTIFIER
     ;