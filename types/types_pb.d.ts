import * as jspb from 'google-protobuf'

import * as google_protobuf_timestamp_pb from 'google-protobuf/google/protobuf/timestamp_pb';


export class DeletedAt extends jspb.Message {
  getTime(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setTime(value?: google_protobuf_timestamp_pb.Timestamp): DeletedAt;
  hasTime(): boolean;
  clearTime(): DeletedAt;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeletedAt.AsObject;
  static toObject(includeInstance: boolean, msg: DeletedAt): DeletedAt.AsObject;
  static serializeBinaryToWriter(message: DeletedAt, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeletedAt;
  static deserializeBinaryFromReader(message: DeletedAt, reader: jspb.BinaryReader): DeletedAt;
}

export namespace DeletedAt {
  export type AsObject = {
    time?: google_protobuf_timestamp_pb.Timestamp.AsObject,
  }
}

