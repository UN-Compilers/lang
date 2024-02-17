parser grammar langParser;


options {
    tokenVocab=langLexer;
}

program: exp* EOF;

var_declaration: IDENTIFIER IDENTIFIER ASSIGN exp;

exp: exp TIMES exp #product
   | exp PLUS exp  #sum
   | LPAR exp RPAR #parens
   | INTEGER       #integer
   | REAL          #real
   | IDENTIFIER    #ident
   ;
