import java.net.*; 
import java.io.*; 

public class HoneycombServer extends Thread { 
    protected Socket clientSocket = null;

    public static void main(String[] args) throws IOException { 
        ServerSocket serverSocket = null; 

        try { 
             serverSocket = new ServerSocket(8888); 
        } 
        catch (IOException e) { 
             System.err.println("Could not listen on port: 8888."); 
             System.exit(1); 
        } 

        try { 
            while (true) {
                System.out.println ("Waiting for new Connection");
                new HoneycombServer(serverSocket.accept()); 
            }
        } 
        catch (IOException e) { 
             System.err.println("Accept failed."); 
             System.exit(1); 
        }

        try {
          serverSocket.close(); 
        } 
        catch (IOException e) { 
          System.err.println("Could not close port: 10008."); 
          System.exit(1); 
        } 
    } 

    private HoneycombServer (Socket clientSoc) {
       clientSocket = clientSoc;
       start();
    }


    public void run() {
      System.out.println ("New Thread Started");

      try { 
           PrintWriter out = new PrintWriter(clientSocket.getOutputStream(), true); 
           BufferedReader in = new BufferedReader(new InputStreamReader( clientSocket.getInputStream())); 

           String inputLine; 

           while ((inputLine = in.readLine()) != null) { 
                System.out.println("Server Received: " + inputLine); 
                out.println("Honeycomb Server has received your input! "); 

                if (inputLine.equals("exit")) 
                    break; 
            } 

           out.close(); 
           in.close(); 
           clientSocket.close(); 
          } 
      catch (IOException e) 
          { 
           System.err.println("Problem with Server");
           System.exit(1); 
          } 
      }
} 