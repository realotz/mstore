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
}

/**
 * @description: Request list return value
 */
export type VolumeResultModel = BasicFetchResult<VolumeListItem>;

export type FileResultModel = BasicFetchResult<FileListItem>;
