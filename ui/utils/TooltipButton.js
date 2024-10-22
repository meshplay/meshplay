import { Button } from '@material-ui/core';
import { CustomTooltip } from '@khulnasoft/sistent';

export default function TooltipButton({ children, onClick, title, variant, ...props }) {
  return (
    <CustomTooltip title={title} placement="top" interactive>
      <Button variant={variant} onClick={onClick} {...props}>
        {children}
      </Button>
    </CustomTooltip>
  );
}
