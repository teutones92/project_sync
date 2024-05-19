import 'package:flutter/material.dart';
import 'package:project_sync/core/config/bloc_config.dart';
import 'package:project_sync/models/app/dashboard_analytic_model/dashboard_analytic_model.dart';

class AnalyticWidget extends StatelessWidget {
  const AnalyticWidget({super.key});

  @override
  Widget build(BuildContext context) {
    return BlocBuilder<DashboardAnalyticBloc, List<DashboardAnalyticModel>>(
      builder: (context, state) {
        return Row(
          children: List.generate(state.length, (index) {
            final item = state[index];
            return Expanded(
              child: Padding(
                padding: EdgeInsets.only(
                    bottom: 20,
                    top: 10,
                    right: index < state.length - 1 ? 20 : 0),
                child: Container(
                  height: 100,
                  decoration: BoxDecoration(
                    color: item.color.withOpacity(0.09),
                    borderRadius: BorderRadius.circular(10),
                  ),
                  child: Center(
                    child: ListTile(
                      leading: Container(
                        width: 50,
                        height: 50,
                        decoration: const BoxDecoration(
                            shape: BoxShape.circle, color: Colors.white),
                        padding: const EdgeInsets.all(10),
                        child: Icon(item.icon, color: item.color, size: 20),
                      ),
                      title: Text(
                        item.value.toString(),
                        style: const TextStyle(fontWeight: FontWeight.bold),
                      ),
                      subtitle: Text(
                        item.subTitle,
                      ),
                    ),
                  ),
                ),
              ),
            );
          }),
        );
      },
    );
  }
}
