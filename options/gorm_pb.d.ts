import * as jspb from 'google-protobuf'

import * as google_protobuf_descriptor_pb from 'google-protobuf/google/protobuf/descriptor_pb';


export class ExtraField extends jspb.Message {
  getType(): string;
  setType(value: string): ExtraField;

  getName(): string;
  setName(value: string): ExtraField;

  getTag(): string;
  setTag(value: string): ExtraField;

  getPackage(): string;
  setPackage(value: string): ExtraField;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ExtraField.AsObject;
  static toObject(includeInstance: boolean, msg: ExtraField): ExtraField.AsObject;
  static serializeBinaryToWriter(message: ExtraField, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ExtraField;
  static deserializeBinaryFromReader(message: ExtraField, reader: jspb.BinaryReader): ExtraField;
}

export namespace ExtraField {
  export type AsObject = {
    type: string,
    name: string,
    tag: string,
    pb_package: string,
  }
}

