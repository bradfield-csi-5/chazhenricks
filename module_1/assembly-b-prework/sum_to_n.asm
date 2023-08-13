section .text
extern _printf
global sum_to_n
sum_to_n:
  mov rax, 0
  cmp rdi, 0
  jg add
	ret

add:
  add rax, rdi
  dec rdi
  cmp rdi, 0
  jg add 
  ret
