import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:project_sync/core/extensions.dart';

import '../../../global/widgets/custom_body_text_field.dart';

class ProjectDeadLine extends StatelessWidget {
  const ProjectDeadLine({
    super.key,
  });

  @override
  Widget build(BuildContext context) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        Text(
          'Project Deadline',
          style: ThemeX.subtitleText,
        ),
        const SizedBox(height: 10),
        Row(
          mainAxisAlignment: MainAxisAlignment.spaceEvenly,
          children: [
            Expanded(
              flex: 4,
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  const Text('Start Date'),
                  CustomBodyTextField(
                      inputFormatters: [FilteringTextInputFormatter.digitsOnly],
                      hintText: 'MM/DD/YYYY',
                      keyboardType: TextInputType.datetime),
                ],
              ),
            ),
            const Spacer(flex: 2),
            Expanded(
              flex: 4,
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  const Text('End Date'),
                  CustomBodyTextField(
                      inputFormatters: [FilteringTextInputFormatter.digitsOnly],
                      hintText: 'MM/DD/YYYY',
                      keyboardType: TextInputType.datetime),
                ],
              ),
            ),
          ],
        ),
      ],
    );
  }
}
