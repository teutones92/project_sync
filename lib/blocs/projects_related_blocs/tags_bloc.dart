import 'package:project_sync/core/config/bloc_config.dart';
import 'package:project_sync/core/config/service_conf.dart';
import 'package:project_sync/models/repository/tags_model/tags_model.dart';

class TagsBloc extends Cubit<List<TagsModel>> {
  TagsBloc() : super([]);

  void addTag(TagsModel tag) {
    state.add(tag);
    emit(state);
  }

  void removeTag(TagsModel tag) {
    state.remove(tag);
    emit(state);
  }

  void clearTags() {
    state.clear();
    emit(state);
  }

  void getTags() async {
    // Fetch tags from the server
    final tags = await TagsService.readTags();

    emit(state);
  }
}
