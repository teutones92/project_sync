class SessionModel {
  final int id;
  final int userId;
  final String token;

  SessionModel({required this.id, required this.userId, required this.token});

  factory SessionModel.fromJson(Map<String, dynamic> json) {
    return SessionModel(
      id: json['id'],
      userId: json['user_id'],
      token: json['token'],
    );
  }
}
