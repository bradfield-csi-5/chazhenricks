section .text
global binary_convert
binary_convert:
  movzx eax, byte [rdi]
  sub eax, '0'
  inc rdi 
  cmp byte [rdi], 00
  jg add_digit
	ret

add_digit:
  shl eax, 1
  movzx esi, byte [rdi] 
  sub esi, '0'
  add eax, esi
  inc rdi 
  cmp byte [rdi], 00
  jg add_digit
	ret

  
