type AvatarProps = {
    src?: string;
}
export const Avatar = (props: AvatarProps) => {
    const {src} = props;
    return <div
        className="rpg-face w-fit"
    >
        <img className="w-32 h-32" src={src} />
    </div>
}
export default Avatar