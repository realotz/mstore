import {
  VolumeParams,
  VolumeResultModel,
  ListFileParams,
  FileResultModel,
  FileReanmeParams,
} from './model/volumeModel';
import { defHttp } from '/@/utils/http/axios';

enum Api {
  Volume = '/v1/volume',
}

export const getVolumeList = (params?: VolumeParams) =>
  defHttp.get<VolumeResultModel>({ url: Api.Volume, params });

export const volumeList = (id: String, params?: ListFileParams) =>
  defHttp.get<FileResultModel>({ url: Api.Volume + '/' + id + '/files', params });

export const fileRename = (id: String, params: FileReanmeParams) =>
  defHttp.post<FileResultModel>({ url: Api.Volume + '/' + id + '/files/rename', params });
