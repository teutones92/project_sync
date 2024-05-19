class StatusCodeModel {
  final int statusCode;
  final String statusCodeMessage;

  StatusCodeModel({required this.statusCode, required this.statusCodeMessage});

  factory StatusCodeModel.fromJson(Map<String, dynamic> json) {
    return StatusCodeModel(
      statusCode: json['status_code'],
      statusCodeMessage: json['status_code_message'],
    );
  }
}
