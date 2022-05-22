import 'package:flutter/material.dart';
import 'package:web/message.dart';
import 'package:web/message_field.dart';
import 'package:web/message_list.dart';

void main() {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Message Board',
      theme: ThemeData(
        primarySwatch: Colors.orange,
      ),
      debugShowCheckedModeBanner: false,
      home: MyHomePage(title: 'AboltuSoft message board'),
    );
  }
}

class MyHomePage extends StatefulWidget {
  List<Message> _messages = exampleMessages;

  MyHomePage({Key? key, required this.title}) : super(key: key);

  final String title;

  @override
  State<MyHomePage> createState() => _MyHomePageState();
}

class _MyHomePageState extends State<MyHomePage> {
  refresh(newMessages) {
    setState(() {
      widget._messages = newMessages;
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text(widget.title),
      ),
      body: Center(
        child: Column(
          children: <Widget>[
            MessageList(messages: widget._messages),
            MessageField(
              notifyParent: refresh,
              messages: widget._messages,
            ),
          ],
        ),
      ),
      floatingActionButtonLocation: FloatingActionButtonLocation.endTop,
      floatingActionButton: FloatingActionButton(
        tooltip: 'Update',
        onPressed: () {},
        child: const Icon(Icons.update_rounded),
      ),
    );
  }
}
