import 'package:flutter/material.dart';
import 'package:project_sync/core/extensions.dart';
import 'package:project_sync/models/repository/tags_model/tags_model.dart';

import '../../../core/config/bloc_config.dart';

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
        const SizedBox(height: 10),
        BlocBuilder<TagsBloc, List<TagsModel>>(
          builder: (context, state) {
            context.read<TagsBloc>().getTags();
            return Wrap(
                spacing: 10,
                children: List.generate(
                  8,
                  (index) => Chip(
                    backgroundColor: ThemeX.of(context).scaffoldBackgroundColor,
                    label: Text(
                      index.toString(),
                      style: const TextStyle(color: Colors.grey),
                    ),
                  ),
                ));
          },
        ),
      ],
    );
  }
}
