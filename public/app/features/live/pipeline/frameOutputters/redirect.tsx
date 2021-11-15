import { RedirectOutputConfig } from '../models.gen';
import { PipelineConfigKind, PipelineFrameOutputterItem } from '../types';

export const redirect: PipelineFrameOutputterItem<RedirectOutputConfig> = {
  kind: PipelineConfigKind.FrameOutputter,
  id: 'redirect',
  description: 'Redirect frame to another channel',
  name: 'Redirect',
  builder: (builder, context) => {
    // builder.addCustomEditor
    console.log('ADD ITEMS!!');
  },
  //viewer: () => <div>JSON AUTO viewer</div>,
};
