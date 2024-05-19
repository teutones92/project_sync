import 'package:flutter/material.dart';

import '../../../core/extensions.dart';
import '../../../global/widgets/custom_body_text_field.dart';

class ProjectName extends StatelessWidget {
  const ProjectName({
    super.key,
  });

  @override
  Widget build(BuildContext context) {
    return Column(
      children: [
        Row(
          children: [
            const Text(
              'Project Name',
              style: TextStyle(fontSize: 20, fontWeight: FontWeight.bold),
            ),
            const SizedBox(width: 10),
            Text('*', style: ThemeX.errorTextStyle)
          ],
        ),
        const SizedBox(height: 10),
        const CustomBodyTextField(hintText: 'Enter Project Name'),
      ],
    );
  }
}
