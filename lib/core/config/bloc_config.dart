import 'package:project_sync/core/config/bloc_config.dart';

export 'package:flutter_bloc/flutter_bloc.dart';
export '../../blocs/auth_related_blocs/login_bloc.dart';
export '../../blocs/dark_mode_bloc.dart';
export '../../blocs/auth_related_blocs/auth_swapper_bloc.dart';
export '../../blocs/auth_related_blocs/register_bloc.dart';
export '../../blocs/home_related_blocs/home_bloc.dart';
export '../../blocs/users_related_blocs/user_data_bloc.dart';
export '../../blocs/home_related_blocs/animate_side_menu_bloc.dart';
export '../../blocs/home_related_blocs/side_menu_bloc.dart';
export '../../blocs/home_related_blocs/dashboard_analytic_bloc.dart';
export '../../blocs/create_project_related_blocs/priority_bloc.dart';
export '../../blocs/create_project_related_blocs/tags_bloc.dart';
export '../../blocs/users_related_blocs/project_lead_bloc.dart';

class MultiBlocsProviderList {
  static get providers => [
        BlocProvider(
          create: (context) => DarkModeBloc(),
        ),
        BlocProvider(
          create: (context) => LoginBloc(),
        ),
        BlocProvider(
          create: (context) => RegisterBloc(),
        ),
        BlocProvider(
          create: (context) => AuthSwapperBloc(),
        ),
        BlocProvider(
          create: (context) => HomeBloc(),
        ),
        BlocProvider(
          create: (context) => UserDataBloc(),
        ),
        BlocProvider(
          create: (context) => AnimateSideMenuBloc(),
        ),
        BlocProvider(
          create: (context) => SideMenuBloc(),
        ),
        BlocProvider(
          create: (context) => DashboardAnalyticBloc(),
        ),
        BlocProvider(
          create: (context) => PriorityBloc(),
        ),
        BlocProvider(
          create: (context) => TagsBloc(),
        ),
        BlocProvider(
          create: (context) => TagSelectBloc(),
        ),
        BlocProvider(
          create: (context) => ProjectLeadBloc(),
        ),
        BlocProvider(
          create: (context) => PrioritySelectedBloc(),
        ),
      ];
}
