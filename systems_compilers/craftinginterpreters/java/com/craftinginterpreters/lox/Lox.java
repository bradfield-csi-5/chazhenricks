package com.craftinginterpreters.lox;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.nio.charset.Charset;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.List;


public class Lox {
    static boolean hadError = false;
    public static void main (String[] args) throws IOException{
        if (args.length > 1){
            System.out.println("Usage: jlox[script]");
            System.exit(64);
        }
        else if (args.length == 1){
            runFile(args[0]);
        }else{
            runPrompt();
        }
    }

    //runs a repl type situation - run jlox with no arguments 
    private static void runPrompt() throws IOException{
        //declare input 
        InputStreamReader input = new InputStreamReader(System.in);
        //pass input into a buffer
        BufferedReader reader = new BufferedReader(input);

        for(;;){
            System.out.print("> ");
            //read in the text as a person types it into the buffered reader
            String line = reader.readLine();
            if (line == null) break;
            //execute the line when we press enter 
            run(line);

            //if user has an error in the REPL, it shouldnt fuck up the whole session
            hadError = false;
        }
    }

    //if we run jlox FILENAME from command line
    private static void runFile(String path) throws IOException {
        //read file into byte string
        byte[] bytes = Files.readAllBytes(Paths.get(path));
        //make a string out of that and run it
        run(new String(bytes, Charset.defaultCharset()));

        //if use has an error in the file execution, exit gracefully
        if(hadError) System.exit(65);
    }


    private static void run(String source) {
        Scanner scanner = new Scanner(source);
        List<Token> tokens = scanner.scanTokens();

        for(Token token : tokens){
            System.out.println(token.toString());
        }
    }

    static void error(int line, String message){
        report(line, "", message);
    }

    private static void report(int line, String where, String message){
        System.err.println("[line " + line + "] Error" + where + ": " + message);
        hadError = true;
    }

}
