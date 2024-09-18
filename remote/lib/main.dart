import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;
import 'dart:convert';

void main() {
  runApp(MyApp());
}

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Remote UI',
      theme: ThemeData(
        primarySwatch: Colors.purple,
        scaffoldBackgroundColor: Colors.grey[900],
      ),
      home: RemoteUIScreen(),
    );
  }
}

class RemoteUIScreen extends StatefulWidget {
  @override
  _RemoteUIScreenState createState() => _RemoteUIScreenState();
}

class _RemoteUIScreenState extends State<RemoteUIScreen> {
  String url = '';

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: SafeArea(
        child: Padding(
          padding: // add padding top and bottom
              EdgeInsets.symmetric(horizontal: 20, vertical: 50),
          child: Column(
            mainAxisAlignment: MainAxisAlignment.spaceBetween,
            children: [
              // Title and Value TextFields
              Column(
                children: [
                  Padding(padding: EdgeInsets.only(top: 20)),
                  Text(
                    'Remote URL',
                    style: TextStyle(
                      fontSize: 20,
                      fontWeight: FontWeight.bold,
                      color: Colors.white,
                    ),
                  ),
                  SizedBox(height: 10),
                  TextField(
                    decoration: InputDecoration(
                      border: OutlineInputBorder(),
                      labelText: 'Enter Remote URL',
                    ),
                    style: TextStyle(color: Colors.white),
                    onChanged: (value) {
                      setState(() {
                        url = value;
                      });
                    },
                  ),
                  SizedBox(height: 20),
                  Icon(Icons.catching_pokemon_rounded),
                ],
              ),

              // Playback and Volume Controls
              Column(
                children: [
                  // Volume Up
                  IconButton(
                    icon: Icon(Icons.arrow_upward_sharp),
                    iconSize: 40,
                    style:
                        IconButton.styleFrom(backgroundColor: Colors.grey[500]),
                    onPressed: () {
                      // TODO: Implement volume up functionality
                      sendCommand('volumeUp');
                    },
                  ),
                  // Playback Controls
                  Row(
                    mainAxisAlignment: MainAxisAlignment.center,
                    children: [
                      IconButton(
                        icon: Icon(Icons.skip_previous_sharp),
                        iconSize: 40,
                        style: IconButton.styleFrom(
                            backgroundColor: Colors.grey[500]),
                        onPressed: () {
                          // TODO: Implement previous track functionality
                          sendCommand('previous');
                        },
                      ),
                      IconButton(
                        icon: Icon(
                          Icons.pause_circle_outline_outlined,
                          color: Colors.grey[500],
                        ),
                        iconSize: 40,
                        padding:
                            EdgeInsets.symmetric(vertical: 20, horizontal: 20),
                        onPressed: () {
                          // TODO: Implement play/pause functionality
                          sendCommand('start');
                        },
                      ),
                      IconButton(
                        icon: Icon(Icons.skip_next_sharp),
                        style: IconButton.styleFrom(
                            backgroundColor: Colors.grey[500]),
                        iconSize: 40,
                        onPressed: () {
                          // TODO: Implement next track functionality
                          sendCommand('next');
                        },
                      ),
                    ],
                  ),
                  // Volume Down
                  IconButton(
                    icon: Icon(Icons.arrow_downward_sharp),
                    style:
                        IconButton.styleFrom(backgroundColor: Colors.grey[500]),
                    iconSize: 40,
                    onPressed: () {
                      // TODO: Implement volume down functionality
                      sendCommand('volumeDown');
                    },
                  ),
                ],
              ),
            ],
          ),
        ),
      ),
    );
  }

  Future<void> sendCommand(String action) async {
    if (url.isEmpty) {
      ScaffoldMessenger.of(context).showSnackBar(
        SnackBar(
          content: Text('Enter a valid URL'),
        ),
      );
      return;
    }

    final command = {'action': action};
    try {
      final response = await http.post(
        Uri.parse(url),
        body: jsonEncode(command),
        headers: {'Content-Type': 'application/json'},
      );

      // Check the status code correctly
      if (response.statusCode >= 200 && response.statusCode < 300) {
        // Status code 200-299 indicates success
        ScaffoldMessenger.of(context).showSnackBar(
          SnackBar(
            content: Text('Command sent successfully'),
          ),
        );
      } else {
        // Status code 400-499 and 500-599 indicate errors
        ScaffoldMessenger.of(context).showSnackBar(
          SnackBar(
            content: Text('Error sending command: ${response.statusCode}'),
          ),
        );
      }
    } catch (e) {
      ScaffoldMessenger.of(context).showSnackBar(
        SnackBar(
          content: Text('Error sending command: $e'),
        ),
      );
    }
  }
}
