import 'package:flutter/material.dart';
import 'package:project_sync/core/extensions.dart';

import 'widgets/contacts_body.dart';

class Contacts extends StatelessWidget {
  const Contacts({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          Text('Contacts', style: ThemeX.titleText),
          const SizedBox(height: 20),
          const Expanded(child: ContactsBody(count: 4)),
        ],
      ),
    );
  }
}
