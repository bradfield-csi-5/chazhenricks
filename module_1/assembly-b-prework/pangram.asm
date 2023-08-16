section .text
global pangram
pangram:
  XOR edx, edx  ;zero out register, this will be our mask 

LOOP_START: ;for each new byte  

  movzx ecx, byte [rdi] ;move the first char into ecx
  cmp ecx, 0x00 ;compare incoming byte to 00
  je END_PROGRAM ;if byte we just moved is null char, then end it 
  or ecx, 32 ;this will force lowercase
  sub ecx, 'a' 
  bts edx, ecx 
  inc rdi 
  jmp LOOP_START



END_PROGRAM:
  xor eax, eax ; set return to 0
  and edx, 0x03ffffff ;will be all 1s if they are the same number 
  cmp edx, 0x03ffffff 
  sete al ;if previous comparison is equal, then this will be set to 0, otherwise 1
  ret

