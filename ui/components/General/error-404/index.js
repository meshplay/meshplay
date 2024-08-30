import { ErrorTypes } from '@/constants/common';
import { useTheme } from '@material-ui/core/styles';
import {
  ErrorSection,
  ErrorSectionContainer,
  ErrorContainer,
  ErrorContentContainer,
  ErrorMain,
  ErrorSectionContent,
  StyledButton,
  ImageContainer,
  IconWrapper,
  Logo,
  LogoText,
  StyledDivider,
  ErrorLink,
} from './styles';
import { Typography, InfoCircleIcon, CustomTooltip } from '@khulnasoft/sistent';
import OrgSwitcher from './OrgSwitcher';
// import RequestForm from './RequestForm';
import CurrentSessionInfo from './CurrentSession';
import { UsesSistent } from '@/components/SistentWrapper';

//TODO: Add component for meshplay version compatiblity error
// const MeshplayVersionCompatiblity = () => {
//   return (
//     <div>
//       <Typography variant="p" component="p" align="center">
//         <InstallMeshplay action={MeshplayAction.UPGRADE.KEY} />
//       </Typography>
//     </div>
//   );
// };

const UnknownServerSideError = (props) => {
  const { errorContent } = props;
  return (
    <div>
      <ErrorContentContainer>
        <Typography variant="p" component="p" align="center">
          {errorContent}
        </Typography>
      </ErrorContentContainer>
    </div>
  );
};

const DefaultError = (props) => {
  const { errorTitle, errorContent, errorType } = props;
  const theme = useTheme();

  return (
    <UsesSistent>
      <ErrorMain>
        <div
          style={{
            display: 'flex',
            flexDirection: 'column',
            alignItems: 'center',
          }}
        >
          <ImageContainer>
            <Logo src="/static/img/meshplay-logo/meshplay-logo.svg" alt="Meshplay logo" />
            <LogoText
              src={
                theme.palette.type === 'dark'
                  ? '/static/img/meshplay-logo/meshplay-white.svg'
                  : '/static/img/meshplay-logo/meshplay-black.svg'
              }
              alt="Meshplay logo text"
            />
          </ImageContainer>
          <Typography variant="h4" component="h4" align="center" className="errormsg">
            {errorTitle
              ? errorTitle
              : "Oops! It seems like you don't have the necessary permissions to view this page."}
          </Typography>
          {errorType === ErrorTypes.UNKNOWN ? (
            <UnknownServerSideError errorContent={errorContent} />
          ) : null}
        </div>
        <ErrorContainer>
          <ErrorSectionContainer>
            <ErrorSection>
              <Typography variant="h5" component="h5" align="center" fontWeight={600}>
                YOUR CURRENT SESSION
              </Typography>
              <CurrentSessionInfo />
            </ErrorSection>
            <StyledDivider orientation="vertical" flexItem />
            <ErrorSection>
              <Typography variant="h5" component="h5" align="center" fontWeight={600}>
                YOUR OPTIONS
              </Typography>
              {/* this is left intentionally inline for now since this is a one off till we implement
               the request form*/}
              <ErrorSectionContent
                style={{
                  flex: '1',
                  justifyContent: 'center',
                }}
              >
                <OrgSwitcher />
                {/*<Divider />
                <RequestForm />*/}
              </ErrorSectionContent>
            </ErrorSection>
          </ErrorSectionContainer>
          <CustomTooltip title="To view the content of this page, switch to an organization where you have more roles using the 'Switch Organization' field.">
            <IconWrapper>
              <InfoCircleIcon height={32} width={32} />
            </IconWrapper>
          </CustomTooltip>
        </ErrorContainer>
        <StyledButton href="/" variant="contained">
          Return to Dashboard
        </StyledButton>
        <Typography variant="textB1Regular" component="p" align="center">
          For more help, please inquire on the
          <ErrorLink href="https://discuss.khulnasoft.com"> discussion forum</ErrorLink> or the{' '}
          <ErrorLink href="https://slack.khulnasoft.com"> Slack workspace</ErrorLink>.
        </Typography>
      </ErrorMain>
    </UsesSistent>
  );
};

export default DefaultError;
