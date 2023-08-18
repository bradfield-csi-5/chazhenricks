section .text
global index
index:
	; rdi: matrix
	; rsi: rows
	; rdx: cols
	; rcx: rindex
	; r8: cindex
  ;Xa + L(C  * i + j)
  ;[rdi] + 4([rdx] * rcx + r8)

  imul rdx, rcx
  add rdx, r8
  shl rdx, 2 
  add rdx, rdi 
  mov rax, [rdx]
	ret
