import java.io.*;
import java.net.*;

public class TestClient {

    public static void main(String[] args) throws IOException {

        String serverHostname = new String ("127.0.0.1");

        if (args.length > 0)
           serverHostname = args[0];

        System.out.println ("Attemping to connect to host " +
		                          serverHostname + " on port 8888.");

        Socket serverSocket = null;
        PrintWriter out = null;
        BufferedReader in = null;

        try {

            serverSocket = new Socket(serverHostname, 8888);
            out = new PrintWriter(serverSocket.getOutputStream(), true); // client input => transfer to server
            in = new BufferedReader(new InputStreamReader(serverSocket.getInputStream())); // server response

        } catch (UnknownHostException e) {
            System.err.println("Don't know about host: " + serverHostname);
            System.exit(1);
        } catch (IOException e) {
            System.err.println("Couldn't get I/O for "
                               + "the connection to: " + serverHostname);
            System.exit(1);
        }

	BufferedReader stdIn = new BufferedReader(new InputStreamReader(System.in)); // read from command line
	String userInput;

    System.out.print ("Please type in your input: ");
	while ((userInput = stdIn.readLine()) != null) {
	    out.println(userInput);
	    System.out.println("Honeycomb Server response: " + in.readLine());
        System.out.print ("input: ");
	}

	out.close();
	in.close();
	stdIn.close();
	serverSocket.close();
    }

}
