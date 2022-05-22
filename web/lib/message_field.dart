import 'package:flutter/material.dart';
import 'package:intl/intl.dart';
import 'package:web/message.dart';
import 'package:web/postmessage.dart';

class MessageField extends StatefulWidget {
  final List<Message> messages;
  final Function notifyParent;

  const MessageField(
      {Key? key, required this.notifyParent, required this.messages})
      : super(key: key);

  @override
  State<StatefulWidget> createState() {
    return _MessageFieldState();
  }
}

class _MessageFieldState extends State<MessageField> {
  final _controller = TextEditingController();
  String _nickname = "";
  String _message = "";

  void _send() {
    setState(() {
      var time = DateTime.now();
      var formatted = DateFormat('dd/MM/yyyy HH:mm').format(time);
      widget.messages.insert(
        0,
        Message(
          author: _nickname,
          message: _message,
          time: formatted,
        ),
      );
      _controller.clear();
      widget.notifyParent(widget.messages);
      createServerMessage(_nickname, _message, formatted);
    });
  }

  @override
  Widget build(BuildContext context) {
    return Container(
      width: MediaQuery.of(context).size.width,
      color: Colors.white60,
      alignment: Alignment.bottomCenter,
      child: LayoutBuilder(
        builder: (BuildContext context, BoxConstraints constraints) {
          return Column(
            children: [
              Row(
                children: [
                  Container(
                    alignment: Alignment.centerLeft,
                    width: constraints.maxWidth,
                    padding:
                        const EdgeInsetsDirectional.fromSTEB(20, 20, 20, 0),
                    child: IntrinsicWidth(
                      child: TextField(
                        keyboardType: TextInputType.multiline,
                        maxLines: 1,
                        decoration: InputDecoration(
                          constraints: const BoxConstraints(minWidth: 94),
                          border: OutlineInputBorder(
                            borderRadius: BorderRadius.circular(15),
                          ),
                          labelText: "nickname",
                        ),
                        onChanged: (text) {
                          _nickname = text;
                        },
                      ),
                    ),
                  ),
                ],
              ),
              Container(
                padding: const EdgeInsets.all(20),
                width: constraints.maxWidth,
                child: TextField(
                  keyboardType: TextInputType.multiline,
                  maxLines: 5,
                  controller: _controller,
                  decoration: InputDecoration(
                    border: OutlineInputBorder(
                      borderRadius: BorderRadius.circular(15),
                    ),
                    labelText: "message",
                    suffixIcon: IconButton(
                      icon: const Icon(Icons.send_rounded),
                      onPressed: _send,
                    ),
                  ),
                  onChanged: (text) {
                    _message = text;
                  },
                ),
              ),
            ],
          );
        },
      ),
    );
  }
}
