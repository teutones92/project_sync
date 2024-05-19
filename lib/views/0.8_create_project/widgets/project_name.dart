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
            Text(
              'Project Name',
              style: ThemeX.subtitleText,
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
