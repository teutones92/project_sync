import 'package:flutter/material.dart';

import '../../../global/widgets/custom_body_text_field.dart';

class ProjectDescription extends StatelessWidget {
  const ProjectDescription({
    super.key,
  });

  @override
  Widget build(BuildContext context) {
    return const Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        Text(
          'Project Description',
          style: TextStyle(fontSize: 20, fontWeight: FontWeight.bold),
        ),
        SizedBox(height: 10),
        CustomBodyTextField(hintText: 'Enter Project Description', maxLines: 5),
      ],
    );
  }
}
