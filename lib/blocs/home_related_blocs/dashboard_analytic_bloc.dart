import 'package:project_sync/core/config/bloc_config.dart';
import 'package:project_sync/models/app/dashboard_analytic_model/dashboard_analytic_model.dart';

class DashboardAnalyticBloc extends Cubit<List<DashboardAnalyticModel>> {
  DashboardAnalyticBloc() : super(DashboardAnalyticModel.data);
}
