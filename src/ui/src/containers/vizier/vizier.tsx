import {Header} from 'components/header/header';
import {SidebarNav} from 'components/sidebar-nav/sidebar-nav';
import * as React from 'react';
import {Route} from 'react-router-dom';
import {AgentDisplay} from './agent-display';
import {DeployInstructions} from './deploy-instructions';
import {QueryManager} from './query-manager';

import './vizier.scss';

// TODO(zasgar/michelle): we should figure out a good way to
// package assets.
// @ts-ignore : TS does not like image files.
import * as infoImage from 'images/icons/agent.svg';
// @ts-ignore : TS does not like image files.
import * as codeImage from 'images/icons/query.svg';
// @ts-ignore : TS does not like image files.
import * as logoImage from 'images/logo.svg';

const PATH_TO_HEADER_TITLE = {
  '/vizier/agents': 'Agents',
  '/vizier/query': 'Query',
};

export interface RouterInfo {
  pathname: string;
}

interface VizierProps {
  match: any;
  location: RouterInfo;
}
export class Vizier extends React.Component<VizierProps, {}> {
  constructor(props) {
    super(props);
  }

  render() {
    const matchPath = this.props.match.path;

    // TODO(michelle): Make an API call to check whether the customer's cluster has already been registered.
    const deployed = true; // Whether or not the cluster has been deployed.
    if (!deployed) {
      return (
        <DeployInstructions
          sitename={window.location.hostname.split('.')[0]}
        />
      );
    }

    return (
      <div className='vizier'>
        <SidebarNav
          logo = {logoImage}
          items={[
            { link: '/vizier/query', selectedImg: codeImage, unselectedImg: codeImage },
            { link: '/vizier/agents', selectedImg: infoImage, unselectedImg: infoImage },
          ]}
        />
        <div className='vizier-body'>
          <Header
             primaryHeading='Pixie Console'
             secondaryHeading={PATH_TO_HEADER_TITLE[this.props.location.pathname]}
          />
          <Route path={`${matchPath}/agents`} component={AgentDisplay} />
          <Route path={`${matchPath}/query`} component={QueryManager} />
        </div>
      </div>
    );
  }
}
export default Vizier;
