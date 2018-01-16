import { connect } from 'react-redux';
import { Dispatch } from 'redux';

import * as actions from '../actions';
import ChartList from '../components/ChartList';
import { StoreState } from '../store/types';

interface RouteProps {
  match: {
    params: {
      repo: string;
    }
  };
}

function mapStateToProps({ charts }: StoreState, { match: { params } }: RouteProps) {
  return {
    charts,
    repo: params.repo
  };
}

function mapDispatchToProps(dispatch: Dispatch<StoreState>) {
  return {
    fetchCharts: (repo: string) => dispatch(actions.fetchCharts(repo))
  };
}

export default connect(mapStateToProps, mapDispatchToProps)(ChartList);
