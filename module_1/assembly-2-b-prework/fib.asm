section .text
global fib

fib: 
  cmp rdi, 1 ; rdi is the argument 1 
  jnbe .recursive
  mov rax, rdi ;if not greater than 1, set return register and return
  ret
.recursive:
  push rdi ;push onto the stack the current state of N
  dec rdi ;go down the n - 1 path
  call fib
  ;will come here when n  - 1 finally returns
  mov rbx, rax ;save the return of n - 1 into a temp variable
  pop rdi ; get n back into rdi 
  sub rdi, 2 ;set rdi up to be n - 2
  push rbx
  call fib
  ;n - 2 will return here
  mov rbp, rax
  pop rbx
  mov rax, rbx
  add rax, rbp
  ret 
