import java.net.*; 
import java.io.*; 

public class HoneycombServer 
{ 
 public static void main(String[] args) throws IOException 
   { 
    ServerSocket serverSocket = null; 

    try { 
         serverSocket = new ServerSocket(8888); 
        } 
    catch (IOException e) 
        { 
         System.err.println("Could not listen on port: 8888."); 
         System.exit(1); 
        } 

    Socket clientSocket = null; 
    System.out.println ("Waiting for connection.....");

    try { 
         clientSocket = serverSocket.accept(); 
        } 
    catch (IOException e) 
        { 
         System.err.println("Accept failed."); 
         System.exit(1); 
        } 

    System.out.println ("Connection successful");
    System.out.println ("Waiting for input.....");

    PrintWriter out = new PrintWriter(clientSocket.getOutputStream(), true); // server response => transfer to client
    BufferedReader in = new BufferedReader(new InputStreamReader(clientSocket.getInputStream())); // client input

    String inputLine; 

    while ((inputLine = in.readLine()) != null) { 
         System.out.println ("Server Received: " + inputLine); 
         out.println("Honeycomb Server has received your input! "); 
    } 

    out.close(); 
    in.close(); 
    clientSocket.close(); 
    serverSocket.close(); 
   } 
} 