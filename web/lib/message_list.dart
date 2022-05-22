import 'package:flutter/material.dart';
import 'package:web/message.dart';

class MessageList extends StatelessWidget {
  final List<Message> messages;

  const MessageList({Key? key, required this.messages}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Expanded(
      child: SingleChildScrollView(
        child: Container(
          padding: const EdgeInsetsDirectional.fromSTEB(0, 0, 0, 20),
          width: MediaQuery.of(context).size.width,
          constraints: const BoxConstraints(minHeight: 450),
          color: Colors.grey,
          child: Column(
            children: messages,
          ),
        ),
      ),
    );
  }
}
