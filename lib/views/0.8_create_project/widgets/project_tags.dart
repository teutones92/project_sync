import 'package:flutter/material.dart';
import 'package:project_sync/core/extensions.dart';

class ProjectTags extends StatelessWidget {
  const ProjectTags({super.key});

  @override
  Widget build(BuildContext context) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        Text(
          'Tags',
          style: ThemeX.subtitleText,
        ),
      ],
    );
  }
}
