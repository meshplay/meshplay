/* eslint-disable no-unused-vars */
import MeshplayOperatorIcon from '../icons/MeshplayOperatorIcon.js';
import fetchMeshplayOperatorStatus from '../../graphql/queries/OperatorStatusQuery';
import ServiceCard from '../ServiceCard';
import { CircularProgress, Grid } from '@material-ui/core';
import MeshplayOperatorDataPanel from '../DataPanels/MeshplayOperator';
import { useEffect, useState } from 'react';
import {
  getOperatorStatusFromQueryResult,
  isMeshplayOperatorConnected,
} from '../helpers/meshplayOperator.js';

// Connection Wizard
// TODO: bind to contextID prop, leaving due to no use in current UI
const MeshplayOperatorScreen = ({ setStepStatus }) => {
  const [operatorInformation, setOperatorInformation] = useState({
    operatorInstalled: false,
    NATSInstalled: false,
    meshSyncInstalled: false,
    operatorVersion: 'N/A',
    meshSyncVersion: 'N/A',
    NATSVersion: 'N/A',
  });
  const [isConnected, setIsConnected] = useState(false);
  const [isLoading, setIsLoading] = useState(false);

  const meshplayOperatorinfo = {
    name: 'Meshplay Operator',
    logoComponent: MeshplayOperatorIcon,
    configComp: <div />,
    operatorInformation,
  };

  const showDataPanel = () => isMeshplayOperatorConnected(operatorInformation);

  useEffect(() => {
    setStepStatus((prev) => ({ ...prev, operator: isConnected }));
  }, [isConnected]);

  useEffect(() => {
    setIsLoading(true);
    fetchMeshplayOperatorStatus().subscribe({
      next: (res) => {
        setIsLoading(false);
        setOperatorState(res);
      },
      error: (err) => setIsLoading(false),
    });
  }, []);

  useEffect(() => {
    setIsConnected(isMeshplayOperatorConnected(operatorInformation));
  }, [operatorInformation]);

  const setOperatorState = (res) => {
    const [isReachable, operatorInformation] = getOperatorStatusFromQueryResult(res);
    setOperatorInformation(operatorInformation);
  };

  return (
    <Grid item xs={12} container justify="center" alignItems="flex-start">
      <Grid
        item
        container
        justify="center"
        alignItems="flex-start"
        lg={6}
        sm={12}
        md={12}
        style={{ paddingLeft: '1rem' }}
      >
        <ServiceCard
          serviceInfo={meshplayOperatorinfo}
          isConnected={isConnected}
          style={{ paddingRight: '1rem' }}
        />
      </Grid>
      <Grid item lg={6} sm={12} md={12} container justify="center">
        {isLoading ? (
          <CircularProgress />
        ) : (
          showDataPanel() && <MeshplayOperatorDataPanel operatorInformation={operatorInformation} />
        )}
      </Grid>
    </Grid>
  );
};

export default MeshplayOperatorScreen;
