import { Tooltip, Fade } from '@material-ui/core';
import {
  DockerIcon,
  GithubIcon,
  MessageIcon,
  SlackIcon,
  SocialContainer,
  SocialMain,
  TwitterIcon,
  YoutubeIcon,
} from './styles';

export default function Socials() {
  return (
    <SocialMain>
      <SocialContainer>
        <Tooltip
          TransitionComponent={Fade}
          TransitionProps={{ timeout: 600 }}
          title="Get connected with the KhulnaSoft community"
        >
          <a href="mailto:community@khulnasoft.com">
            <MessageIcon height={45} width={45} />
          </a>
        </Tooltip>

        <Tooltip
          TransitionComponent={Fade}
          TransitionProps={{ timeout: 600 }}
          title="Join the community Slack"
        >
          <a href="https://slack.khulnasoft.com">
            <SlackIcon height={45} width={45} />
          </a>
        </Tooltip>

        <Tooltip
          TransitionComponent={Fade}
          TransitionProps={{ timeout: 600 }}
          title="Follow KhulnaSoft on Twitter"
        >
          <a href="https://twitter.com/khulnasoft">
            <TwitterIcon height={40} width={40} />
          </a>
        </Tooltip>

        <Tooltip
          TransitionComponent={Fade}
          TransitionProps={{ timeout: 600 }}
          title="Contribute to KhulnaSoft projects"
        >
          <a href="https://github.com/khulnasoft">
            <GithubIcon height={45} width={45} />
          </a>
        </Tooltip>

        <Tooltip
          TransitionComponent={Fade}
          TransitionProps={{ timeout: 600 }}
          title="Watch community meeting recordings"
        >
          <a href="https://www.youtube.com/playlist?list=PL3A-A6hPO2IMPPqVjuzgqNU5xwnFFn3n0">
            <YoutubeIcon height={45} width={45} />
          </a>
        </Tooltip>

        <Tooltip
          TransitionComponent={Fade}
          TransitionProps={{ timeout: 600 }}
          title="Access Docker images"
        >
          <a href="https://hub.docker.com/u/khulnasoft/">
            <DockerIcon height={45} width={45} />
          </a>
        </Tooltip>
      </SocialContainer>
    </SocialMain>
  );
}
