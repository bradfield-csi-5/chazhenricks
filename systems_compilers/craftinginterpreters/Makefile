# Compile the Java interpreter .java files to .class files.
generate_ast: 
	@ $(MAKE) -f util/java.make DIR=java PACKAGE=tool
	@ java -cp build/java com.craftinginterpreters.tool.GenerateAst \
		java/com/craftinginterpreters/lox
		
jlox: generate_ast
	@ $(MAKE) -f util/java.make DIR=java PACKAGE=lox

