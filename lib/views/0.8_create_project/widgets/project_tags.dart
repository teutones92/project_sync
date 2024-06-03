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
        FutureBuilder(
            future: context.read<TagsBloc>().getTags(),
            builder: (context, snapshot) {
              if (!snapshot.hasData) {
                return const Center(child: CircularProgressIndicator());
              }

              return BlocBuilder<TagsBloc, List<TagsModel>>(
                builder: (context, state) {
                  return BlocBuilder<TagSelectBloc, int?>(
                      builder: (context, selectedIndex) {
                    return Wrap(
                      children: [
                        Wrap(
                          children: List.generate(state.length, (index) {
                            final item = state[index];
                            return Card(
                              elevation: 8,
                              child: InkWell(
                                onTap: () {
                                  context
                                      .read<TagsBloc>()
                                      .selectTag(item, index);
                                  context
                                      .read<TagSelectBloc>()
                                      .selectTag(index);
                                  if (item.id == null &&
                                      index != selectedIndex) {
                                    // Show dialog to add new tag
                                    _showDialogMethod(context);
                                  }
                                },
                                child: Chip(
                                  backgroundColor: selectedIndex == index
                                      ? ThemeX.goldColor.shade800
                                      : ThemeX.of(context)
                                          .scaffoldBackgroundColor,
                                  label: Text(
                                    item.tagName,
                                    style: TextStyle(
                                      color: ThemeX.of(context).brightness ==
                                              Brightness.dark
                                          ? Colors.white
                                          : Colors.grey.shade800,
                                    ),
                                  ),
                                ),
                              ),
                            );
                          }),
                        ),
                      ],
                    );
                  });
                },
              );
            }),
      ],
    );
  }

  Future<dynamic> _showDialogMethod(BuildContext context) {
    return showDialog(
      context: context,
      builder: (context) {
        final controller = TextEditingController();
        return AlertDialog(
          title: const Text('Add New Tag'),
          content: TextField(
            controller: controller,
            textCapitalization: TextCapitalization.words,
            decoration: const InputDecoration(
              hintText: 'Enter tag name',
            ),
          ),
          actions: [
            TextButton(
              onPressed: () {
                Navigator.pop(context);
                context.read<TagSelectBloc>().clearTag();
              },
              child: const Text('Cancel'),
            ),
            TextButton(
              onPressed: controller.text.isNotEmpty
                  ? () {
                      context.read<TagsBloc>().addTag(TagsModel(
                            id: null,
                            tagName: controller.text,
                          ));
                      Navigator.pop(context);
                    }
                  : null,
              child: const Text('Add'),
            ),
          ],
        );
      },
    );
  }
}
