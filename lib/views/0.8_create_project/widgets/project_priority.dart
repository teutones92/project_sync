import 'package:flutter/material.dart';
import 'package:project_sync/models/repository/priority_model/priority_model.dart';

import '../../../core/config/bloc_config.dart';
import '../../../core/extensions.dart';

class ProjectPriority extends StatelessWidget {
  const ProjectPriority({super.key});

  @override
  Widget build(BuildContext context) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        Text(
          'Priority',
          style: ThemeX.subtitleText,
        ),
        const SizedBox(height: 10),
        BlocBuilder<PriorityBloc, List<PriorityModel>>(
          builder: (context, state) {
            context.read<PriorityBloc>().getPriorities(context);
            return Container();
          },
        )
      ],
    );
  }
}
