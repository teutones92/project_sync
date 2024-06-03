import 'package:project_sync/core/config/bloc_config.dart';
import 'package:project_sync/core/config/service_conf.dart';
import 'package:project_sync/models/repository/tags_model/tags_model.dart';

class TagSelectBloc extends Cubit<int?> {
  TagSelectBloc() : super(null);

  void selectTag(int index) {
    if (index == state) {
      emit(null);
    } else {
      emit(index);
    }
  }

  void clearTag() {
    emit(null);
  }
}

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

  void selectTag(TagsModel tag, int index) {
    state[index] = tag;
    emit(List.from(state));
  }

  Future<List<TagsModel>> getTags() async {
    // Fetch tags from the server
    final tags = await TagsService.readTags();
    if (tags is List<TagsModel>) {
      state.clear();
      state.addAll(tags);
      state.add(TagsModel(id: null, tagName: 'Other'));
    }
    return tags;
  }
}
