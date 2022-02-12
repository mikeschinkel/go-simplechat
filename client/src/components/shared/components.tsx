import { Container } from '@mui/material';


interface IPageWrapperProps {
    children: string | JSX.Element | JSX.Element[];
}


/**
 * Standard container for major URL pages.
 * 
 * @param props
 * @returns 
 */
export function PageWrapper(props: IPageWrapperProps): JSX.Element {
    return (
        <Container maxWidth='lg'>
            {props.children}
        </Container>
    );
}
