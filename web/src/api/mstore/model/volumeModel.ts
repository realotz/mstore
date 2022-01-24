import { BasicPageParams, BasicFetchResult } from '/@/api/model/baseModel';

export type VolumeParams = {
  account?: string;
  nickname?: string;
  option?: BasicPageParams;
};

export type ListFileParams = {
  path?: string;
  type?: number;
  option?: BasicPageParams;
};

export type FileReanmeParams = {
  path: string;
  new_path: string;
};
export type FileParams = {
  id: string;
  path: string;
};

export type DelFileParams = {
  files: FileParams[];
};

export type CopyMoveParams = {
  files: FileParams[];
  is_delete: boolean;
  to_path: string;
  to_volume_id: string;
};

export interface VolumeListItem {
  id: string;
  name: string;
  provider: string;
  created_at: number;
  updated_at: number;
}

export interface FileListItem {
  name: string;
  provider: string;
  string: string;
  path: string;
  is_dir: boolean;
  size: number;
  updated_at: number;
  volume_id: string;
}

/**
 * @description: Request list return value
 */
export type VolumeResultModel = BasicFetchResult<VolumeListItem>;

export type FileResultModel = BasicFetchResult<FileListItem>;
