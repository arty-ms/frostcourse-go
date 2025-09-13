export default function ProgressBar({ current, total }) {
    const progress = (current / total) * 100;

    return (
        <div className="bg-tg-secondary-bg h-2 rounded-full my-5 overflow-hidden">
            <div
                className="bg-tg-button h-full transition-all duration-300 ease-out rounded-full"
                style={{ width: `${progress}%` }}
            />
        </div>
    );
}