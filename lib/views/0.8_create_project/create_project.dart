import 'package:flutter/material.dart';
import 'package:project_sync/core/extensions.dart';
import 'package:project_sync/global/widgets/header_widget_items.dart';

import '../../core/config/project_conf_widgets.dart';

class CreateProject extends StatelessWidget {
  const CreateProject({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        elevation: 0,
        backgroundColor: Colors.transparent,
        toolbarHeight: 90,
        title: const HeaderWidgetItems(),
      ),
      floatingActionButton: FloatingActionButton.extended(
        onPressed: () {
          NavigatorX.pop(context);
        },
        label: const Text('Save Project'),
        icon: const Icon(Icons.save_alt_rounded),
      ),
      body: Padding(
        padding: const EdgeInsets.all(40),
        child: ListView(
          children: [
            Text(
              'Create New Project',
              style: ThemeX.titleText,
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
          ],
        ),
      ),
    );
  }
}
