import 'package:flutter/material.dart';
import 'package:project_sync/core/config/bloc_config.dart';
import 'package:project_sync/core/extensions.dart';
import 'package:project_sync/global/widgets/header_widget_items.dart';

import '../../core/config/project_conf_widgets.dart';

class CreateProject extends StatelessWidget {
  const CreateProject({super.key});

  @override
  Widget build(BuildContext context) {
    context.read<ProjectLeadBloc>().setProjectLead(
          context.read<UserDataBloc>().state,
          context,
        );
    return SafeArea(
      child: Scaffold(
        floatingActionButton: FloatingActionButton.extended(
          onPressed: () {
            NavigatorX.pop(context);
          },
          label: const Text('Save Project'),
          icon: const Icon(Icons.save_alt_rounded),
        ),
        body: CustomScrollView(
          slivers: [
            SliverAppBar(
              backgroundColor: ThemeX.of(context).scaffoldBackgroundColor,
              title: const HeaderWidgetItems(),
            ),
            SliverToBoxAdapter(
              child: Padding(
                padding: const EdgeInsets.all(40),
                child: Column(
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: [
                    Text(
                      'Create Project',
                      style: ThemeX.titleText,
                    ),
                    TextButton(
                      onPressed: () {
                        // Show dialog to add new tag
                      },
                      style: TextButton.styleFrom(padding: EdgeInsets.zero),
                      child: Tooltip(
                        message: 'By default, you are the project lead.',
                        child: BlocBuilder<ProjectLeadBloc, UserState?>(
                          builder: (context, state) {
                            return Text(
                              'Project Lead: ${state!.username}',
                            );
                          },
                        ),
                      ),
                    ),
                    const SizedBox(height: 20),
                    const ProjectName(),
                    const SizedBox(height: 20),
                    const ProjectDescription(),
                    const SizedBox(height: 20),
                    const ProjectDeadLine(),
                    const SizedBox(height: 20),
                    const ProjectTeamMembers(),
                    const SizedBox(height: 20),
                    const ProjectTags(),
                    const SizedBox(height: 20),
                    const ProjectPriority(),
                    const SizedBox(height: 40),
                  ],
                ),
              ),
            ),
          ],
        ),
      ),
    );
  }
}
