import type { AppRouteModule } from '/@/router/types';

import { LAYOUT } from '/@/router/constant';
import { t } from '/@/hooks/web/useI18n';

const volume: AppRouteModule = {
  path: '/volume',
  name: 'Volume',
  component: LAYOUT,
  redirect: '/volume/index',
  meta: {
    orderNo: 10,
    hideChildrenInMenu: true,
    icon: 'ion:grid-outline',
    title: t('routes.volume.title'),
  },
  children: [
    {
      path: 'index',
      name: 'VolumeIndex',
      component: () => import('/@/views/volume/index.vue'),
      meta: {
        title: t('routes.volume.index'),
        icon: 'simple-icons:about-dot-me',
        hideMenu: true,
      },
    },
  ],
};

export default volume;
