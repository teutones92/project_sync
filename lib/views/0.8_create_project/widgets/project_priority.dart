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
        FutureBuilder(
            future: context.read<PriorityBloc>().getPriorities(context),
            builder: (context, snapshot) {
              if (!snapshot.hasData) {
                return const Center(child: CircularProgressIndicator());
              }
              return BlocBuilder<PriorityBloc, List<PriorityModel>>(
                builder: (context, state) {
                  return BlocBuilder<PrioritySelectedBloc, int?>(
                    builder: (context, selectedState) {
                      return Wrap(
                        children: List.generate(state.length, (index) {
                          final item = state[index];
                          return Card(
                            elevation: 8,
                            child: InkWell(
                              onTap: () {
                                context
                                    .read<PrioritySelectedBloc>()
                                    .selectPriority(index, item);
                              },
                              child: Chip(
                                backgroundColor: index == selectedState
                                    ? ThemeX.goldColor.shade800
                                    : ThemeX.of(context)
                                        .scaffoldBackgroundColor,
                                shape: RoundedRectangleBorder(
                                  borderRadius: BorderRadius.circular(5),
                                ),
                                label: Text(
                                  item.name,
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
                      );
                    },
                  );
                },
              );
            })
      ],
    );
  }
}
