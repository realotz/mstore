import {
  VolumeParams,
  VolumeResultModel,
  ListFileParams,
  FileResultModel,
  FileReanmeParams,
  CopyMoveParams,
  DelFileParams,
  CreateFileParams,
  FileParams,
  FileDownResultModel,
  SaveFileParams,
} from './model/volumeModel';
import { defHttp } from '/@/utils/http/axios';
import { ErrorMessageMode } from '/#/axios';
enum Api {
  Volume = '/v1/volume',
}

export const getVolumeList = (params?: VolumeParams) =>
  defHttp.get<VolumeResultModel>({ url: Api.Volume, params });

export const volumeList = (id: String, params?: ListFileParams) =>
  defHttp.get<FileResultModel>({ url: Api.Volume + '/' + id + '/files', params });

export const fileRename = (id: String, params: FileReanmeParams) =>
  defHttp.post<FileResultModel>({ url: Api.Volume + '/' + id + '/files/rename', params });

export const createFile = (id: String, params: CreateFileParams) =>
  defHttp.post<FileResultModel>({ url: Api.Volume + '/' + id + '/files', params });

export const copyMove = (params: CopyMoveParams, mode: ErrorMessageMode = 'modal') =>
  defHttp.post<FileResultModel>(
    { url: Api.Volume + '/files/copy-move', params },
    {
      errorMessageMode: mode,
    },
  );

export const saveFile = (id: String, params: SaveFileParams) =>
  defHttp.put<FileResultModel>({ url: Api.Volume + '/' + id + '/files', params });

export const delFile = (params: DelFileParams) =>
  defHttp.post<FileResultModel>({ url: Api.Volume + '/files/del', params: params });

export const downUrl = (id: String, params: FileParams) =>
  defHttp.get<FileDownResultModel>({ url: Api.Volume + '/' + id + '/files/down', params });
