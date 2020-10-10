
#include <stdio.h>
#include<string.h>
#include<math.h>
#include<stdlib.h>
#define BLANK ' '
#define TAB '\t'
#define MAX 50
void push(long int symbol);
long int pop();
void infix_to_postfix();
long int eval_post();
int priority(char symbol);
int isEmpty();
int white_space(char symbol);
char infix[MAX], postfix[MAX];
long int stack[MAX];
int top;

int main(){
long int value;
top=-1;
printf("Enter infix : ");
gets(infix);
infix_to_postfix();
printf("Postfix : %s\n",postfix);
value = eval_post();
printf("Value of expression : %ld\n", value);
return 0;
} /*End of main()*/
void infix_to_postfix()
{
    unsigned int i, p = 0;
    char next;
    char symbol;
    for (i = 0; i < strlen(infix); i++)
    {
        symbol = infix[i];
        if (!white_space(symbol))
        {
            switch (symbol)
            {
            case '(':
                push(symbol);
                break;
            case ')':
                while ((next = pop() )!= '(')
                    postfix[p++] = next;
                break;
            case '+':
            case '-':
            case '*':
            }
          
        }        
     }