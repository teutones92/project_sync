import 'package:flutter/material.dart';
import 'package:project_sync/views/0.2_dashboard/widgets/analytic_widget/analytic_widget.dart';

import 'widgets/my_task_view_widget/my_task_view_widget.dart';
import 'widgets/projects_view_widget/projects_view_widget.dart';

class Dashboard extends StatelessWidget {
  const Dashboard({super.key});

  @override
  Widget build(BuildContext context) {
    return const Column(
      children: [
        AnalyticWidget(),
        Expanded(
          child: Row(
            children: [
              ProjectsViewWidget(),
              MyTaskViewWidget(),
            ],
          ),
        )
      ],
    );
  }
}
